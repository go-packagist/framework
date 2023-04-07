package bootstraps

import "github.com/go-packagist/framework/foundation"

type Bootstrapper interface {
	Bootstrap(*foundation.Application)
}
