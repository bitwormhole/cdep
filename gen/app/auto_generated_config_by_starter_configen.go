// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package app

import (
	dep0xc9214e "github.com/bitwormhole/cdep/dep"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
    
)


func nop(x ... interface{}){
	util.Int64ToTime(0)
	lang.CreateReleasePool()
}


func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()
	nop(err,cominfobuilder)

	// component: com0-dep0xc9214e.App
	cominfobuilder.Next()
	cominfobuilder.ID("com0-dep0xc9214e.App").Class("life").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComApp{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComApp : the factory of component: com0-dep0xc9214e.App
type comFactory4pComApp struct {

    mPrototype * dep0xc9214e.App

	

}

func (inst * comFactory4pComApp) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComApp) newObject() * dep0xc9214e.App {
	return & dep0xc9214e.App {}
}

func (inst * comFactory4pComApp) castObject(instance application.ComponentInstance) * dep0xc9214e.App {
	return instance.Get().(*dep0xc9214e.App)
}

func (inst * comFactory4pComApp) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComApp) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComApp) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComApp) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComApp) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComApp) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}




