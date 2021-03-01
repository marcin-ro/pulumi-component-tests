package main

import (
	"fmt"

	"github.com/pulumi/pulumi-random/sdk/v2/go/random"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(createResources)
}

var appNames = []string{"appA", "appB"}

func createResources(ctx *pulumi.Context) error {
	for _, appName := range appNames {
		component, err := newAppComponent(ctx, appName)
		if err != nil {
			return err
		}
		for i := 0; i < 2; i++ {
			resName := fmt.Sprintf("pass-%d-%s", i, appName)
			_, err := random.NewRandomPassword(
				ctx,
				resName,
				&random.RandomPasswordArgs{Length: pulumi.Int(10)},
				pulumi.Parent(component),
			)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func newAppComponent(ctx *pulumi.Context, appName string) (*AppComponent, error) {
	appComponent := AppComponent{}
	name := fmt.Sprintf("component-%s", appName)
	return &appComponent, ctx.RegisterComponentResource("marcin:main/appComponent:AppComponent", name, &appComponent)
}

type AppComponent struct {
	pulumi.ResourceState
}
