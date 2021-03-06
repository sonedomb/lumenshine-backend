package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/Soneso/lumenshine-backend/pb"

	"github.com/Soneso/lumenshine-backend/admin/client"
	mw "github.com/Soneso/lumenshine-backend/admin/middleware"
	cerr "github.com/Soneso/lumenshine-backend/icop_error"

	"github.com/Soneso/lumenshine-backend/constants"
	"github.com/Soneso/lumenshine-backend/helpers"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/gin-gonic/gin"

	"github.com/Soneso/lumenshine-backend/admin/config"
	"github.com/Soneso/lumenshine-backend/admin/db"
	"github.com/Soneso/lumenshine-backend/admin/route"
	tt "github.com/Soneso/lumenshine-backend/admin/templates"
	"github.com/Soneso/lumenshine-backend/db/pageinate"
	qq "github.com/Soneso/lumenshine-backend/db/querying"

	"strconv"

	m "github.com/Soneso/lumenshine-backend/services/db/models"
)

const (
	//CustomerRoutePrefix is the prefix for the customer group. We need this in order to get all the routes for this base url
	CustomerRoutePrefix = "customer"
)

//init setup all the routes for the users handling
func init() {
	route.AddRoute("GET", "/list", CustomerList, []string{}, "customer_list", CustomerRoutePrefix)
	route.AddRoute("GET", "/details/:id", CustomerDetails, []string{}, "customer_details", CustomerRoutePrefix)
	route.AddRoute("POST", "/update_personal_data", CustomerEdit, []string{}, "customer_update_personal_data", CustomerRoutePrefix)
	route.AddRoute("GET", "/orders/:id", CustomerOrders, []string{}, "customer_orders", CustomerRoutePrefix)
	route.AddRoute("GET", "/wallets/:id", CustomerWallets, []string{}, "customer_wallets", CustomerRoutePrefix)
	route.AddRoute("POST", "/update_kyc_status", CustomerUpdateKYCStatus, []string{}, "update_kyc_status", CustomerRoutePrefix)
	route.AddRoute("POST", "/reset2fa", Reset2fa, []string{}, "customer_reset2fa", CustomerRoutePrefix)
}

//AddCustomerRoutes adds all the routes for the user handling
func AddCustomerRoutes(rg *gin.RouterGroup) {
	for _, r := range route.GetRoutesForPrefix(CustomerRoutePrefix) {
		f := r.HandlerFunc.(func(uc *mw.AdminContext, c *gin.Context))
		rg.Handle(r.Method, r.Prefix+r.Path, mw.UseAdminContext(f, r.Name))
	}
}

//CustomerListRequest for filtering the customers
//swagger:parameters CustomerListRequest CustomerList
type CustomerListRequest struct {
	pageinate.PaginationRequestStruct

	FilterCustomerID int      `form:"filter_customer_id" json:"filter_customer_id"`
	FilterForeName   string   `form:"filter_forename" json:"filter_forename"`
	FilterLastName   string   `form:"filter_lastname" json:"filter_lastname"`
	FilterEmail      string   `form:"filter_email" json:"filter_email"`
	FilterKycStatus  []string `form:"filter_kyc_status" json:"filter_kyc_status"`

	SortCustomerID       string `form:"sort_customer_id" json:"sort_customer_id"`
	SortForeName         string `form:"sort_forename" json:"sort_forename"`
	SortLastName         string `form:"sort_lastname" json:"sort_lastname"`
	SortEmail            string `form:"sort_email" json:"sort_email"`
	SortRegistrationDate string `form:"sort_registration_date" json:"sort_registration_date"`
	SortLastLogin        string `form:"sort_last_login" json:"sort_last_login"`
}

//CustomerListItem is one item in the list
// swagger:model
type CustomerListItem struct {
	ID               int       `json:"id"`
	Forename         string    `json:"forename"`
	Lastname         string    `json:"last_name"`
	Email            string    `json:"email"`
	KycStatus        string    `json:"kyc_status"`
	RegistrationDate time.Time `json:"registration_date"`
	LastLogin        time.Time `json:"last_login"`
}

//CustomerListResponse list of customers
// swagger:model
type CustomerListResponse struct {
	pageinate.PaginationResponseStruct
	Items []CustomerListItem `json:"items"`
}

//CustomerList returns list of all customers, filtered by given params
// swagger:route GET /portal/admin/dash/customer/list customer CustomerList
//
// Returns list of all customers, filtered by given params
//
// Consumes:
//     - multipart/form-data
//
//     Produces:
//     - application/json
//
//     Responses:
//       200:CustomerListResponse
func CustomerList(uc *mw.AdminContext, c *gin.Context) {
	var err error
	var rr CustomerListRequest
	if err = c.Bind(&rr); err != nil {
		c.JSON(http.StatusBadRequest, cerr.LogAndReturnError(uc.Log, err, cerr.ValidBadInputData, cerr.BindError))
		return
	}

	if valid, validErrors := cerr.ValidateStruct(uc.Log, rr); !valid {
		c.JSON(http.StatusBadRequest, validErrors)
		return
	}

	//this is the initial queryMod
	//we will append queries and sorting to it
	q := []qm.QueryMod{
		qm.Select(
			m.UserProfileColumns.ID,
			m.UserProfileColumns.Forename,
			m.UserProfileColumns.Lastname,
			m.UserProfileColumns.Email,
			m.UserProfileColumns.KycStatus,
			m.UserProfileColumns.CreatedAt,
		),
	}

	if rr.FilterCustomerID != 0 {
		q = append(q, qm.Where("id=?", rr.FilterCustomerID))
	}

	if rr.FilterForeName != "" {
		q = append(q, qm.Where("forename ilike ?", qq.Like(rr.FilterForeName)))
	}

	if rr.FilterLastName != "" {
		q = append(q, qm.Where("lastname ilike ?", qq.Like(rr.FilterLastName)))
	}

	if rr.FilterEmail != "" {
		q = append(q, qm.Where("email ilike ? ", qq.Like(rr.FilterEmail)))
	}

	if len(rr.FilterKycStatus) > 0 {
		filter := make([]interface{}, len(rr.FilterKycStatus))
		for i := range rr.FilterKycStatus {
			filter[i] = rr.FilterKycStatus[i]
		}
		q = append(q, qm.WhereIn("kyc_status in ? ", filter...))
	}

	r := new(CustomerListResponse)

	//we need to get the total count before sorting and applying the pagination
	r.TotalCount, err = m.UserProfiles(q...).Count(db.DBC)
	if err != nil {
		c.JSON(http.StatusInternalServerError, cerr.LogAndReturnError(uc.Log, err, "Error getting total count", cerr.GeneralError))
		return
	}

	//apply sorting
	q = qq.AddSorting(rr.SortCustomerID, m.UserProfileColumns.ID, q)
	q = qq.AddSorting(rr.SortForeName, m.UserProfileColumns.Forename, q)
	q = qq.AddSorting(rr.SortLastName, m.UserProfileColumns.Lastname, q)
	q = qq.AddSorting(rr.SortEmail, m.UserProfileColumns.Email, q)
	q = qq.AddSorting(rr.SortRegistrationDate, m.UserProfileColumns.CreatedAt, q)

	qP := pageinate.Paginate(q, &rr.PaginationRequestStruct, &r.PaginationResponseStruct)
	customers, err := m.UserProfiles(qP...).All(db.DBC)
	if err != nil && err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, cerr.LogAndReturnError(uc.Log, err, "Error getting users", cerr.GeneralError))
		return
	}

	r.Items = make([]CustomerListItem, len(customers))
	for i, c := range customers {
		r.Items[i] = CustomerListItem{
			ID:               c.ID,
			Forename:         c.Forename,
			Lastname:         c.Lastname,
			Email:            c.Email,
			KycStatus:        c.KycStatus,
			RegistrationDate: c.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, r)
}

// CustomerDetailsResponse - customer details response
// swagger:model
type CustomerDetailsResponse struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	Forename   string `json:"forename"`
	Lastname   string `json:"lastname"`
	Salutation string `json:"salutation"`
	//LastLogin        time.Time `json:"last_login"`
	Address           string     `json:"address"`
	ZipCode           string     `json:"zip_code"`
	City              string     `json:"city"`
	State             string     `json:"state"`
	CountryCode       string     `json:"country_code"`
	Nationality       string     `json:"nationality"`
	MobileNR          string     `json:"mobile_nr"`
	BirthDay          *time.Time `json:"birth_day"`
	BirthPlace        string     `json:"birth_place"`
	AdditionalName    string     `json:"additional_name"`
	BirthCountryCode  string     `json:"birth_country_code"`
	BankAccountNumber string     `json:"bank_account_number"`
	BankNumber        string     `json:"bank_number"`
	BankPhoneNumber   string     `json:"bank_phone_number"`
	TaxID             string     `json:"tax_id"`
	TaxIDName         string     `json:"tax_id_name"`
	OccupationName    string     `json:"occupation_name"`
	OccupationCode08  string     `json:"occupation_code08"`
	OccupationCode88  string     `json:"occupation_code88"`
	EmployerName      string     `json:"employer_name"`
	EmployerAddress   string     `json:"employer_address"`
	LanguageCode      string     `json:"language_code"`
	RegistrationDate  time.Time  `json:"registration_date"`
}

//CustomerDetails returns details of specified customer
// swagger:route GET /portal/admin/dash/customer/details/:id customer CustomerDetails
//
// Returns details of specified customer
//
// Consumes:
//     - multipart/form-data
//
//     Produces:
//     - application/json
//
//     Responses:
//       200:CustomerDetailsResponse
func CustomerDetails(uc *mw.AdminContext, c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, cerr.LogAndReturnError(uc.Log, err, "Error parsing id", cerr.GeneralError))
		return
	}

	u, err := m.UserProfiles(
		qm.Where("id=?", id),
	).One(db.DBC)
	if err != nil && err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, cerr.LogAndReturnError(uc.Log, err, "Error getting user from db", cerr.GeneralError))
		return
	}

	if err == sql.ErrNoRows {
		c.JSON(http.StatusBadRequest, cerr.NewIcopError("id", cerr.UserNotExists, "User does not exist in db", ""))
		return
	}
	response := CustomerDetailsResponse{
		ID:                u.ID,
		Email:             u.Email,
		Forename:          u.Forename,
		Lastname:          u.Lastname,
		Salutation:        u.Salutation,
		Address:           u.Address,
		ZipCode:           u.ZipCode,
		City:              u.City,
		State:             u.State,
		CountryCode:       u.CountryCode,
		Nationality:       u.Nationality,
		MobileNR:          u.MobileNR,
		BirthPlace:        u.BirthPlace,
		AdditionalName:    u.AdditionalName,
		BirthCountryCode:  u.BirthCountryCode,
		BankAccountNumber: u.BankAccountNumber,
		BankNumber:        u.BankNumber,
		BankPhoneNumber:   u.BankPhoneNumber,
		TaxID:             u.TaxID,
		TaxIDName:         u.TaxIDName,
		OccupationName:    u.OccupationName,
		OccupationCode08:  u.OccupationCode08,
		OccupationCode88:  u.OccupationCode88,
		EmployerName:      u.EmployerName,
		EmployerAddress:   u.EmployerAddress,
		LanguageCode:      u.LanguageCode,
		RegistrationDate:  u.CreatedAt,
	}
	if !u.BirthDay.IsZero() {
		response.BirthDay = &u.BirthDay
	}
	c.JSON(http.StatusOK, &response)
}

// CustomerEditRequest - request data
//swagger:parameters CustomerEditRequest CustomerEdit
type CustomerEditRequest struct {
	//required : true
	ID                int    `form:"id" json:"id"`
	Forename          string `form:"forename" json:"forename" validate:"required,icop_nonum,min_trim=2,max=64"`
	Lastname          string `form:"lastname" json:"lastname" validate:"required,icop_nonum,min_trim=2,max=64"`
	Salutation        string `form:"salutation" json:"salutation" validate:"max=64"`
	Address           string `form:"address" json:"address" validate:"max=512"`
	ZipCode           string `form:"zip_code" json:"zip_code" validate:"max=32"`
	City              string `form:"city" json:"city" validate:"max=128"`
	State             string `form:"state" json:"state" validate:"max=128"`
	CountryCode       string `form:"country_code" json:"country_code" validate:"max=2"`
	Nationality       string `form:"nationality" json:"nationality" validate:"max=128"`
	MobileNR          string `form:"mobile_nr" json:"mobile_nr" validate:"max=64"`
	BirthDay          string `form:"birth_day" json:"birth_day"`
	BirthPlace        string `form:"birth_place" json:"birth_place" validate:"max=128"`
	AdditionalName    string `form:"additional_name" json:"additional_name" validate:"max=255"`
	BirthCountryCode  string `form:"birth_country_code" json:"birth_country_code" validate:"max=3"`
	BankAccountNumber string `form:"bank_account_number" json:"bank_account_number" validate:"max=255"`
	BankNumber        string `form:"bank_number" json:"bank_number" validate:"max=255"`
	BankPhoneNumber   string `form:"bank_phone_number" json:"bank_phone_number" validate:"max=255"`
	TaxID             string `form:"tax_id" json:"tax_id" validate:"max=255"`
	TaxIDName         string `form:"tax_id_name" json:"tax_id_name" validate:"max=255"`
	OccupationName    string `form:"occupation_name" json:"occupation_name" validate:"max=256"`
	OccupationCode08  string `form:"occupation_code08" json:"occupation_code08" validate:"max=8"`
	OccupationCode88  string `form:"occupation_code88" json:"occupation_code88" validate:"max=8"`
	EmployerName      string `form:"employer_name" json:"employer_name" validate:"max=500"`
	EmployerAddress   string `form:"employer_address" json:"employer_address" validate:"max=500"`
	LanguageCode      string `form:"language_code" json:"language_code" validate:"max=10"`
}

//CustomerEdit updates customer details and returns customer
// swagger:route POST /portal/admin/dash/customer/update_personal_data customer CustomerEdit
//
// Updates customer details and returns customer
//
// Consumes:
//     - multipart/form-data
//
//     Produces:
//     - application/json
//
//     Responses:
//       200:CustomerDetailsResponse
func CustomerEdit(uc *mw.AdminContext, c *gin.Context) {
	var err error
	var rr CustomerEditRequest
	if err = c.Bind(&rr); err != nil {
		c.JSON(http.StatusBadRequest, cerr.LogAndReturnError(uc.Log, err, cerr.ValidBadInputData, cerr.BindError))
		return
	}

	if valid, validErrors := cerr.ValidateStruct(uc.Log, rr); !valid {
		c.JSON(http.StatusBadRequest, validErrors)
		return
	}

	//get the birthday
	birthDay, err := time.Parse("2006-01-02", rr.BirthDay)
	if rr.BirthDay != "" && err != nil {
		c.JSON(http.StatusBadRequest, cerr.NewIcopError("birth_day", cerr.InvalidArgument, "Birthday wrong format", ""))
		return
	}

	u, err := m.UserProfiles(
		qm.Where("id=?", rr.ID),
	).One(db.DBC)
	if err != nil && err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, cerr.LogAndReturnError(uc.Log, err, "Error getting user from db", cerr.GeneralError))
		return
	}

	if err == sql.ErrNoRows {
		c.JSON(http.StatusBadRequest, cerr.NewIcopError("id", cerr.UserNotExists, "User does not exist in db", ""))
		return
	}

	u.Forename = rr.Forename
	u.Lastname = rr.Lastname
	u.Salutation = rr.Salutation
	u.Address = rr.Address
	u.ZipCode = rr.ZipCode
	u.City = rr.City
	u.State = rr.State
	u.CountryCode = rr.CountryCode
	u.Nationality = rr.Nationality
	u.MobileNR = rr.MobileNR
	u.BirthDay = birthDay
	u.BirthPlace = rr.BirthPlace

	u.AdditionalName = rr.AdditionalName
	u.BirthCountryCode = rr.BirthCountryCode
	u.BankAccountNumber = rr.BankAccountNumber
	u.BankNumber = rr.BankNumber
	u.BankPhoneNumber = rr.BankPhoneNumber
	u.TaxID = rr.TaxID
	u.TaxIDName = rr.TaxIDName
	u.OccupationName = rr.OccupationName
	u.OccupationCode08 = rr.OccupationCode08
	u.OccupationCode88 = rr.OccupationCode88
	u.EmployerName = rr.EmployerName
	u.EmployerAddress = rr.EmployerAddress
	u.LanguageCode = rr.LanguageCode

	u.UpdatedBy = getUpdatedBy(c)
	u.UpdatedAt = time.Now().In(boil.GetLocation())

	_, err = u.Update(db.DBC, boil.Infer())

	if err != nil {
		c.JSON(http.StatusInternalServerError, cerr.LogAndReturnError(uc.Log, err, "Error updating user", cerr.GeneralError))
		return
	}

	response := CustomerDetailsResponse{
		ID:                u.ID,
		Forename:          u.Forename,
		Lastname:          u.Lastname,
		Email:             u.Email,
		Salutation:        u.Salutation,
		Address:           u.Address,
		ZipCode:           u.ZipCode,
		City:              u.City,
		State:             u.State,
		CountryCode:       u.CountryCode,
		Nationality:       u.Nationality,
		MobileNR:          u.MobileNR,
		BirthPlace:        u.BirthPlace,
		AdditionalName:    u.AdditionalName,
		BirthCountryCode:  u.BirthCountryCode,
		BankAccountNumber: u.BankAccountNumber,
		BankNumber:        u.BankNumber,
		BankPhoneNumber:   u.BankPhoneNumber,
		TaxID:             u.TaxID,
		TaxIDName:         u.TaxIDName,
		OccupationName:    u.OccupationName,
		OccupationCode08:  u.OccupationCode08,
		OccupationCode88:  u.OccupationCode88,
		EmployerName:      u.EmployerName,
		EmployerAddress:   u.EmployerAddress,
		LanguageCode:      u.LanguageCode,
		RegistrationDate:  u.CreatedAt,
	}
	if !u.BirthDay.IsZero() {
		response.BirthDay = &u.BirthDay
	}
	c.JSON(http.StatusOK, &response)
}

//CustomerOrdersRequest to get the orders
//swagger:parameters CustomerOrdersRequest CustomerOrders
type CustomerOrdersRequest struct {
	pageinate.PaginationRequestStruct
}

//OrderListItem is one item in the list
// swagger:model
type OrderListItem struct {
	ID     int       `json:"id"`
	Date   time.Time `json:"date"`
	Amount int64     `json:"amount"`
	Price  float64   `json:"price"`
	Chain  string    `json:"chain"`
	Status string    `json:"status"`
}

//OrderListResponse list of orders
// swagger:model
type OrderListResponse struct {
	pageinate.PaginationResponseStruct
	Items []OrderListItem `json:"items"`
}

//CustomerOrders returns list of all customer's orders
// swagger:route GET /portal/admin/dash/customer/orders customer CustomerOrders
//
// Returns list of all customer's orders
//
// Consumes:
//     - multipart/form-data
//
//     Produces:
//     - application/json
//
//     Responses:
//       200:OrderListResponse
func CustomerOrders(uc *mw.AdminContext, c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, cerr.LogAndReturnError(uc.Log, err, "Error parsing id", cerr.GeneralError))
		return
	}

	var rr CustomerOrdersRequest
	if err = c.Bind(&rr); err != nil {
		c.JSON(http.StatusBadRequest, cerr.LogAndReturnError(uc.Log, err, cerr.ValidBadInputData, cerr.BindError))
		return
	}

	if valid, validErrors := cerr.ValidateStruct(uc.Log, rr); !valid {
		c.JSON(http.StatusBadRequest, validErrors)
		return
	}

	q := []qm.QueryMod{
		qm.Select(
			m.UserOrderColumns.ID,
			m.UserOrderColumns.TokenAmount,
			m.UserOrderColumns.ExchangeCurrencyDenominationAmount,
			m.UserOrderColumns.PaymentNetwork,
			m.UserOrderColumns.OrderStatus,
			m.UserOrderColumns.CreatedAt,
		),
	}
	q = append(q, qm.Where(m.UserOrderColumns.UserID+"=?", id))

	r := new(OrderListResponse)

	//we need to get the total count before sorting and applying the pagination
	r.TotalCount, err = m.UserOrders(q...).Count(db.DBC)
	if err != nil {
		c.JSON(http.StatusInternalServerError, cerr.LogAndReturnError(uc.Log, err, "Error getting total count", cerr.GeneralError))
		return
	}

	q = append(q, qm.OrderBy(m.UserOrderColumns.CreatedAt))

	qP := pageinate.Paginate(q, &rr.PaginationRequestStruct, &r.PaginationResponseStruct)
	orders, err := m.UserOrders(qP...).All(db.DBC)
	if err != nil && err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, cerr.LogAndReturnError(uc.Log, err, "Error getting orders", cerr.GeneralError))
		return
	}

	r.Items = make([]OrderListItem, len(orders))
	/*for i, o := range orders {
		v, ok := o.ChainAmount.Float64()
		if !ok {
			v = 0
		}
		r.Items[i] = OrderListItem{
			ID:     o.ID,
			Date:   o.CreatedAt,
			Amount: o.TokenAmount,
			Price:  v,
			Chain:  o.Chain,
			Status: o.OrderStatus,
		}
	}*/
	c.JSON(http.StatusOK, r)
}

//CustomerWalletsRequest to get the wallets
//swagger:parameters CustomerWalletsRequest CustomerWallets
type CustomerWalletsRequest struct {
	pageinate.PaginationRequestStruct
}

//WalletListItem is one item in the list
// swagger:model
type WalletListItem struct {
	Name       string `json:"name"`
	PublicKey  string `json:"public_key"`
	WalletType string `json:"wallet_type"`
}

//WalletListResponse list of wallets
// swagger:model
type WalletListResponse struct {
	pageinate.PaginationResponseStruct
	Items []WalletListItem `json:"items"`
}

//CustomerWallets returns list of all customer's wallets
// swagger:route GET /portal/admin/dash/customer/wallets/:id customer CustomerWallets
//
// Returns list of all customer's wallets
//
// Consumes:
//     - multipart/form-data
//
//     Produces:
//     - application/json
//
//     Responses:
//       200:WalletListResponse
func CustomerWallets(uc *mw.AdminContext, c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, cerr.LogAndReturnError(uc.Log, err, "Error parsing id", cerr.GeneralError))
		return
	}

	var rr CustomerWalletsRequest
	if err = c.Bind(&rr); err != nil {
		c.JSON(http.StatusBadRequest, cerr.LogAndReturnError(uc.Log, err, cerr.ValidBadInputData, cerr.BindError))
		return
	}

	if valid, validErrors := cerr.ValidateStruct(uc.Log, rr); !valid {
		c.JSON(http.StatusBadRequest, validErrors)
		return
	}

	q := []qm.QueryMod{
		qm.Select(
			m.UserWalletColumns.WalletName,
			m.UserWalletColumns.PublicKey,
			m.UserWalletColumns.WalletType,
		),
	}
	q = append(q, qm.Where(m.UserWalletColumns.UserID+"=?", id))

	r := new(WalletListResponse)

	//we need to get the total count before sorting and applying the pagination
	r.TotalCount, err = m.UserWallets(q...).Count(db.DBC)
	if err != nil {
		c.JSON(http.StatusInternalServerError, cerr.LogAndReturnError(uc.Log, err, "Error getting total count", cerr.GeneralError))
		return
	}

	q = append(q, qm.OrderBy(m.UserWalletColumns.OrderNR))

	qP := pageinate.Paginate(q, &rr.PaginationRequestStruct, &r.PaginationResponseStruct)
	wallets, err := m.UserWallets(qP...).All(db.DBC)
	if err != nil && err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, cerr.LogAndReturnError(uc.Log, err, "Error getting wallets", cerr.GeneralError))
		return
	}

	r.Items = make([]WalletListItem, len(wallets))
	for i, w := range wallets {
		r.Items[i] = WalletListItem{
			Name:       w.WalletName,
			PublicKey:  w.PublicKey,
			WalletType: w.WalletType,
		}
	}
	c.JSON(http.StatusOK, r)
}

//CustomerUpdateKYCStatusRequest - request
//swagger:parameters CustomerUpdateKYCStatusRequest CustomerUpdateKYCStatus
type CustomerUpdateKYCStatusRequest struct {
	// the customer id
	//required : true
	ID int `form:"id" json:"id"`
	//KYC status, e.g. InReview, Pending
	//required : true
	KycStatus string `form:"kyc_status" json:"kyc_status" validate:"required,max=64"`
}

//CustomerUpdateKYCStatus updates status
// swagger:route POST /portal/admin/dash/customer/update_kyc_status customer CustomerUpdateKYCStatus
//
// Updates KYC status
//
// Consumes:
//     - multipart/form-data
//
//     Produces:
//     - application/json
//
//     Responses:
//       200:
func CustomerUpdateKYCStatus(uc *mw.AdminContext, c *gin.Context) {
	var err error
	var rr CustomerUpdateKYCStatusRequest
	if err = c.Bind(&rr); err != nil {
		c.JSON(http.StatusBadRequest, cerr.LogAndReturnError(uc.Log, err, cerr.ValidBadInputData, cerr.BindError))
		return
	}

	if valid, validErrors := cerr.ValidateStruct(uc.Log, rr); !valid {
		c.JSON(http.StatusBadRequest, validErrors)
		return
	}

	if rr.KycStatus != m.KycStatusApproved &&
		rr.KycStatus != m.KycStatusInReview &&
		rr.KycStatus != m.KycStatusNotSupported &&
		rr.KycStatus != m.KycStatusPending &&
		rr.KycStatus != m.KycStatusRejected &&
		rr.KycStatus != m.KycStatusWaitingForData &&
		rr.KycStatus != m.KycStatusWaitingForReview {
		c.JSON(http.StatusBadRequest, cerr.NewIcopError("kyc_status", cerr.InvalidArgument, "Invalid status value", ""))
		return
	}

	u, err := m.UserProfiles(
		qm.Where("id=?", rr.ID),
		qm.Select(
			m.UserProfileColumns.ID,
			m.UserProfileColumns.KycStatus,
		),
	).One(db.DBC)
	if err != nil && err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, cerr.LogAndReturnError(uc.Log, err, "Error getting user from db", cerr.GeneralError))
		return
	}
	if err == sql.ErrNoRows {
		c.JSON(http.StatusBadRequest, cerr.NewIcopError("id", cerr.UserNotExists, "User does not exist in db", ""))
		return
	}

	u.KycStatus = rr.KycStatus
	u.UpdatedBy = getUpdatedBy(c)
	u.UpdatedAt = time.Now().In(boil.GetLocation())

	_, err = u.Update(db.DBC, boil.Whitelist(m.UserProfileColumns.ID,
		m.UserProfileColumns.KycStatus,
		m.UserProfileColumns.UpdatedBy,
		m.UserProfileColumns.UpdatedAt))

	if err != nil {
		c.JSON(http.StatusInternalServerError, cerr.LogAndReturnError(uc.Log, err, "Error updating user", cerr.GeneralError))
		return
	}

	c.JSON(http.StatusOK, "{}")
}

//Reset2faRequest - request
//swagger:parameters Reset2faRequest Reset2fa
type Reset2faRequest struct {
	// the customer id
	//required : true
	ID int `form:"id" json:"id"`
}

//Reset2fa resets the flag and sends the email
// swagger:route POST /portal/admin/dash/customer/reset2fa customer Reset2fa
//
// Resets the flag and sends the email
//
// Consumes:
//     - multipart/form-data
//
//     Produces:
//     - application/json
//
//     Responses:
//       200:
func Reset2fa(uc *mw.AdminContext, c *gin.Context) {
	var err error
	var rr Reset2faRequest
	if err = c.Bind(&rr); err != nil {
		c.JSON(http.StatusBadRequest, cerr.LogAndReturnError(uc.Log, err, cerr.ValidBadInputData, cerr.BindError))
		return
	}

	u, err := m.UserProfiles(
		qm.Where("id=?", rr.ID),
		qm.Select(
			m.UserProfileColumns.ID,
			m.UserProfileColumns.Email,
			m.UserProfileColumns.Forename,
			m.UserProfileColumns.Lastname,
			m.UserProfileColumns.Reset2faByAdmin,
		),
	).One(db.DBC)
	if err != nil && err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, cerr.LogAndReturnError(uc.Log, err, "Error getting user from db", cerr.GeneralError))
		return
	}

	if err == sql.ErrNoRows {
		c.JSON(http.StatusBadRequest, cerr.NewIcopError("id", cerr.UserNotExists, "User does not exist in db", ""))
		return
	}

	u.Reset2faByAdmin = true
	u.MailConfirmationKey = helpers.RandomString(constants.DefaultMailkeyLength)
	u.MailConfirmationExpiryDate = time.Unix(time.Now().AddDate(0, 0, constants.DefaultMailkeyExpiryDays).Unix(), 0)
	u.UpdatedBy = getUpdatedBy(c)
	u.UpdatedAt = time.Now().In(boil.GetLocation())

	_, err = u.Update(db.DBC, boil.Whitelist(m.UserProfileColumns.ID,
		m.UserProfileColumns.Reset2faByAdmin,
		m.UserProfileColumns.MailConfirmationKey,
		m.UserProfileColumns.MailConfirmationExpiryDate,
		m.UserProfileColumns.UpdatedBy,
		m.UserProfileColumns.UpdatedAt))

	if err != nil {
		c.JSON(http.StatusInternalServerError, cerr.LogAndReturnError(uc.Log, err, "Error updating user", cerr.GeneralError))
		return
	}

	langCode := "en"
	msgSubject := fmt.Sprintf("%s :: Your new 2FA Secret", config.Cnf.Site.SiteName)
	msgBody := tt.RenderTemplateToString(uc, c, "reset_tfa_mail", langCode, gin.H{
		"TokenUrl":  config.Cnf.WebLinks.LostTFA + u.MailConfirmationKey,
		"ImagesUrl": config.Cnf.WebLinks.ImagesUrl,
		"TokenValidTo": helpers.TimeToString(
			u.MailConfirmationExpiryDate, langCode,
		),
	})

	_, err = client.MailClient.SendMail(c, &pb.SendMailRequest{
		Base:    &pb.BaseRequest{RequestId: uc.RequestID, UpdateBy: getUpdatedBy(c)},
		From:    config.Cnf.Site.EmailSender,
		To:      u.Email,
		Subject: msgSubject,
		Body:    msgBody,
		IsHtml:  true,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, cerr.LogAndReturnError(uc.Log, err, "Error sending mail to user", cerr.GeneralError))
		return
	}

	c.JSON(http.StatusOK, "{}")
}
