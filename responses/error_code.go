package responses

const (
	FullNameTidakBolehKosong                              = 4000
	ParameterBodyTidakSesuai                              = 4001
	FormatTokenTidakBenar                                 = 4002
	SilahkanLoginTerlebihDahulu                           = 4003
	DeviceIDTidakSesuai                                   = 4004
	HeaderDeviceAuthorizationBukanTokenDeviceAutorization = 4005
	TokenBasicAuthTidakBolehKosong                        = 4006
	TokenBasicAuthTidakBenar                              = 4007
	ErrorKetikaMendapatkanDataUser                        = 4008
	TokenTidakBolehKosong                                 = 4009
	IdTidakBoleh0                                         = 4010
	InvalidToken                                          = 1111
	UserNotAllowedAccess                                  = 1113
	TokenIsNotAllowedEmpty                                = 1114
	TokenExpired                                          = 1115
)

var ErrorCodeText = map[int]string{
	FullNameTidakBolehKosong:                              "Fullname tidak boleh kosong",
	ParameterBodyTidakSesuai:                              "Parameter body tidak sesuai",
	FormatTokenTidakBenar:                                 "Format token tidak benar",
	SilahkanLoginTerlebihDahulu:                           "Silahkan login terlebih dahulu",
	DeviceIDTidakSesuai:                                   "Device ID tidak sesuai",
	HeaderDeviceAuthorizationBukanTokenDeviceAutorization: "Header DeviceAuthorization bukan token device autorization",
	TokenBasicAuthTidakBolehKosong:                        "Token basic auth tidak boleh kosong",
	TokenBasicAuthTidakBenar:                              "Token basic auth tidak benar",
	ErrorKetikaMendapatkanDataUser:                        "Error ketika mendapatkan data user",
	TokenTidakBolehKosong:                                 "Token tidak boleh kosong",
	IdTidakBoleh0:                                         "Id tidak boleh 0",
	InvalidToken:                                          "Invalid Token!",
	UserNotAllowedAccess:                                  "User not allowed access",
	TokenIsNotAllowedEmpty:                                "Token is not allowed empty",
	TokenExpired:                                          "Token expired",
}
