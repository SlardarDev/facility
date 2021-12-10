package parallel_loader

import (
	"context"
	"fmt"
)

type TLoader struct {
	H string
}

func (c TLoader) Load(ctx context.Context, ctxParam interface{}) error {
	fmt.Println(c.H)
	return nil
}

func (c TLoader) Name() string {
	return c.H
}

func ExampleLoad() {

	loaderMgr := LoadManager{
		LoaderContainer{
			LoaderContainer{TLoader{"1"}, TLoader{"2"}}, TLoader{"3"},
		},
		LoaderContainer{
			LoaderContainer{TLoader{"5"}, TLoader{"6"}, LoaderContainer{TLoader{"8"}, TLoader{"9"}}, TLoader{"10"}}, TLoader{"7"},
		},
	}
	loaderMgr.Load(context.Background(), nil)
	//Unordered output:
	//1
	//2
	//3
	//5
	//6
	//7
	//8
	//9
	//10
}
