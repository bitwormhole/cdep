package dep

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

// App 是应用的主体
type App struct {
	markup.Component `class:"life"`
}

func (inst *App) _Impl() (application.LifeRegistry, application.Looper) {
	return inst, inst
}

// GetLifeRegistration 获取需要注册的life
func (inst *App) GetLifeRegistration() *application.LifeRegistration {
	lr := &application.LifeRegistration{}
	lr.Looper = inst
	return lr
}

// Loop 运行主循环
func (inst *App) Loop() error {
	return nil
}
