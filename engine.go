package vendingMachine


type vendingMachine struct{
	insertedMoney int
	coins map[string]int
	items map[string]int
}

func NewVendingMachine(coins, items map[string]int) *vendingMachine {
	return &vendingMachine{coins: coins, items: items}
}

func (m vendingMachine) InsertedMoney() int {
	return m.insertedMoney
}

func (m *vendingMachine) InsertCoin(coin string) {
	m.insertedMoney += m.coins[coin]
}

func (m *vendingMachine) SelectSD() string {
	price := 18
	change := m.insertedMoney - price
	return "SD" + m.change(change)
}
func (m *vendingMachine) SelectCC() string {
	price := 12
	change := m.insertedMoney - price
	//m.insertedMoney = 0
	return "CC" + m.change(change)
}
func  (m *vendingMachine) change(c int) string {
	var str string
	values := [...]int{10, 5, 2, 1}
	coins := [...]string{"T", "F", "TW", "O"}

	for i := 0 ; i < len(values); i++ {
		if c >= values[i] {
			str += ", " + coins[i]
			c -= values[i]
		}
	}
	return str
}
func (vm *vendingMachine) CoinReturn() string {
	coins := vm.change(vm.insertedMoney)
	vm.insertedMoney = 0
	return coins[2:len(coins)]
}