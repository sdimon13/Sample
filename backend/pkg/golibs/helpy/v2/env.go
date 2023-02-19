package helpy

const (
	EnvLocal      = "local"
	EnvQa         = "qa"
	EnvDev        = "dev"
	EnvProd       = "prod"
	EnvDemo       = "demo"
	EnvDemoHG     = "demohg"
	EnvStand      = "stand"
	EnvKgz        = "kgz"
	EnvDevKgz     = "devkgz"
	EnvBoaipProd  = "boaipprod"
	EnvFm         = "fm"
	EnvFmProd     = "fmprod"
	EnvGeoProd    = "geoprod"
	EnvKz         = "kz"
	GroupEnvRao   = "rao"
	GroupEnvKgz   = "kgz"
	GroupEnvBoaip = "boaip"
	GroupEnvKz    = "kz"
	GroupEnvGeo   = "geo"
	GroupEnvFm    = "fm"
)

func IsRegional(group, env string) bool {
	return group == GetRegionalGroup(env)
}

func IsProduction(env string) bool {
	switch env {
	case EnvProd, EnvKgz, EnvBoaipProd, EnvFmProd, EnvGeoProd, EnvKz:
		return true
	}
	return false
}

func GetRegionalGroup(env string) string {
	switch env {
	// RAO
	case EnvDev:
		fallthrough
	case EnvProd:
		fallthrough
	case EnvLocal:
		fallthrough
	case EnvStand:
		fallthrough
	case EnvDemo:
		fallthrough
	case EnvDemoHG:
		fallthrough
	case EnvQa:
		return GroupEnvRao
	// Киргизия
	case EnvDevKgz:
		fallthrough
	case EnvKgz:
		return GroupEnvKgz
	// ФМ
	case EnvFm:
		fallthrough
	case EnvFmProd:
		return GroupEnvFm
	// Казахстан
	case EnvKz:
		return GroupEnvKz
	// Белорусь
	case EnvBoaipProd:
		return GroupEnvBoaip
	// Грузия
	case EnvGeoProd:
		return GroupEnvGeo
	}
	return ""
}
