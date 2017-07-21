// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	"reflect"
	"github.com/revel/revel"
	controllers0 "github.com/revel/modules/static/app/controllers"
	_ "github.com/revel/modules/testrunner/app"
	controllers1 "github.com/revel/modules/testrunner/app/controllers"
	_ "ilo/app"
	controllers "ilo/app/controllers"
	models "ilo/app/models"
	tests "ilo/tests"
	"github.com/revel/revel/testing"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.INFO.Println("Running revel server")
	
	revel.RegisterController((*controllers.GorpController)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Begin",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Commit",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Rollback",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					76: []string{ 
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Suite",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					129: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Monitor)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					17: []string{ 
						"ilos",
						"systems",
					},
				},
			},
			&revel.MethodType{
				Name: "AddiLOForm",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					21: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "AddiLO",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "ilo", Type: reflect.TypeOf((*models.Ilo)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Overview",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "ilo_id", Type: reflect.TypeOf((*int64)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					34: []string{ 
						"ilo_id",
						"totalHealthMap",
					},
				},
			},
			&revel.MethodType{
				Name: "Fans",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "ilo_id", Type: reflect.TypeOf((*int64)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					40: []string{ 
						"ilo_id",
						"fans",
					},
				},
			},
			&revel.MethodType{
				Name: "Powers",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "ilo_id", Type: reflect.TypeOf((*int64)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					45: []string{ 
						"ilo_id",
						"powers",
					},
				},
			},
			&revel.MethodType{
				Name: "Temperatures",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "ilo_id", Type: reflect.TypeOf((*int64)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					50: []string{ 
						"ilo_id",
						"temperatures",
					},
				},
			},
			&revel.MethodType{
				Name: "EventLog",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "ilo_id", Type: reflect.TypeOf((*int64)(nil)) },
					&revel.MethodArg{Name: "pageNumber", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					66: []string{ 
						"ilo_id",
					},
				},
			},
			&revel.MethodType{
				Name: "SystemLog",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "ilo_id", Type: reflect.TypeOf((*int64)(nil)) },
					&revel.MethodArg{Name: "pageNumber", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					82: []string{ 
						"ilo_id",
					},
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
	}
	testing.TestSuites = []interface{}{ 
		(*tests.AppTest)(nil),
	}

	revel.Run(*port)
}
