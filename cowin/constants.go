package cowin

const (
	base_url = "https://cdn-api.co-vin.in/api/v2"

	states_list_url    = base_url + "/admin/location/states"
	districts_list_url = base_url + "/admin/location/districts"

	availability_by_pin_code_url = base_url + "/appointment/sessions/public/calendarByPin"
	availability_by_district_url = base_url + "/appointment/sessions/public/calendarByDistrict"
)
