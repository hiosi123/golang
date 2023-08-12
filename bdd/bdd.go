package bdd

// import (
// 	"testing"

// 	. "github.com/onsi/ginkgo"
// 	. "github.com/onsi/gomega"
// )

// // 1. BDD (behaviour driven development is a testing methodology that has evolved from Test Driven Development.)

// //Ginkgo and Gomega

// //Gomega is an assertion library => best paired with the Ginkgo BDD test framework

// //Describe: expressivley organise the test code.

// //Context: contatiner desribes a scenario.

// //It: closure defines a certain spc that we want to assert

// //Expect: performs the actual assertion, we are going to use gomega for this.

// func TestAllSuite(t *testing.T) {
// 	RegisterFailHandler(Fail)
// 	RunSpecs(t, "example suite")
// }

// var _ = Describe("Checking the Add function", func() { //행동

// 	BeforeSuite(func() {
// 		//initialize repository connect db
// 	})

// 	Describe("Fetching all exisiting vidoes", func() {

// 		Context("Valid inputs for add function", func() { //시나리오

// 			BeforeEach(func() {
// 				//save something in repo db
// 			})

// 			It("Should return the correct result given 2 integers", func() {
// 				result, _ := Add(2, 2)
// 				OmegaMatcher(result).Should(Equal(4)) //assertion
// 			})

// 			It("should do something", func() {

// 			})

// 			AfterEach(func() {
// 				// delete something in repo
// 			})
// 		})
// 	})

// 	AfterSuite(func() {
// 		//close db
// 	})

// })
