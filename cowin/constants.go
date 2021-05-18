package cowin

const (
	baseUrl = "https://cdn-api.co-vin.in/api/v2"

	districtsListUrl = baseUrl + "/admin/location/districts"

	availabilityByPinCodeUrl  = baseUrl + "/appointment/sessions/public/calendarByPin"
	availabilityByDistrictUrl = baseUrl + "/appointment/sessions/public/calendarByDistrict"
)
