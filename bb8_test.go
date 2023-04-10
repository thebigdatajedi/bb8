package main

import (
	//"github.com/stretchr/testify/assert"
	"go/types"
	"testing"
)

// First::
// Create a function in the test file and start the name with Test.
// The function should take *testing.T from the GoLang testing package.
// The test function signature is Test<UnitUnderTestName>(t *testing.T) with <UnitUnderTestName> as the name
// of the function we are testing.
func TestUsage(t *testing.T) {
	//Second:: Usage() is the function doesn't have any args, so I'm commenting them out instead of deleting them.
	//to keep track of the features of the test function.
	//type args struct { // function args definition
	//	a int
	//	b int
	//}

	//Third:: Keeping in mind the::
	//The Three Parts of a Test
	//1. Arrange
	//2. Act
	//3. Assert

	//I'll be looking at this next part of the test.
	_ = []struct { // test case definition
		name string
		//args args // removed args because Usage() doesn't have any args
		want types.Nil // looks like want is expected result
	}{ // test cases
		{ // test case
			//looks like above we declare the elements of the test case struct, it's name, args and want.
			//in this case Usage returns nothing so we are testing for that.
			//In logical branching what we are doing is testing the other right side of the
			//if statement, essentially.
			name: "should return nothing",
			//args: args{10, 20},
			want: types.Nil{},
		},
	} //defining the struct by declaring the variables name, args, want & the test cases where the
	//variables are set is the Arrange works in Go.

	//Fourth:: The Act part of the test.
	//tt is an item in the range.
	//for _, tt := range tests { // test cases executions
	//t is the *testing.T we passed into the function at the very beginning.
	//t.Run(tt.name, func(t *testing.T) {
	//	//_, err := Usage()
	//	//if err != nil {
	//	//	//Fifth:: The Assert part of the test. - in this case I'm checkin that it returns nil.
	//	//	assert.Fail(t, "Usage() error = %v", err)
	//	//	return
	//	}
	//})
}

//Final notes:: All my functions are side effect functions so I'm not returning anything.
//I'm just logging the errors and returning nothing.
//I'm going to have to create integration tests for these functions and I don't know how to do it just yet.
//I'll have to look into it.
