package test

//test package for qunit
import QUnit "github.com/rusco/qunit"
import "strconv"

//fictive TestScenario
type Scenario struct{}

func (s Scenario) Setup() {
	print("Hi, I am the Setup Function")
}
func (s Scenario) Teardown() {
	print("Hi, I am the Teardown Function")
}

func main() {

	QUnit.ModuleLifecycle("A", Scenario{})
	QUnit.Test("just a test", func(assert QUnit.QUnitAssert) {
		QUnit.Expect(1)
		assert.Ok(true, "")
	})

	QUnit.Module("B")
	QUnit.Test("test 1", func(assert QUnit.QUnitAssert) {
		square := func(x int) int {
			return x * x
		}
		result := square(2)
		assert.Equal(strconv.Itoa(result), strconv.Itoa(4), "square(2) equals 4")
	})
	QUnit.Test("test 2", func(assert QUnit.QUnitAssert) {
		assert.Ok(true, "true succeeds")
	})

	QUnit.Module("C")
	QUnit.Test("test 3", func(assert QUnit.QUnitAssert) {
		assert.Ok(true, "0 means false")
	})
}
