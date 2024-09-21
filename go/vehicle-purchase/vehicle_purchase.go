package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle.
// Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection.
// It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	result := option1
	if result > option2 {
		result = option2
	}
	return result + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice * 0.80
	} else if age >= 10 {
		return originalPrice * 0.50
	} else {
		return originalPrice * 0.70
	}
}

// 3. Рассчитайте ориентировочную стоимость подержанного автомобиля
// Теперь, когда вы приняли решение, вы хотите убедиться,
// что получите справедливую цену в дилерском центре.
// Поскольку вы заинтересованы в покупке подержанного автомобиля,
// цена зависит от возраста транспортного средства.
// Для приблизительной оценки предположим,
// что если автомобилю меньше 3 лет, его стоимость составляет 80%
// от первоначальной цены, которая была у него, когда он был совершенно новым.
// Если ему не менее 10 лет, его стоимость составляет 50%.
// Если автомобилю не менее 3, но менее 10 лет, его стоимость составляет 70% от первоначальной цены.

// Реализуйте CalculateResellPrice(originalPrice, age) функцию, которая применяет эту логику, используя if, else if и else (есть и другие способы, если вы хотите попрактиковаться). В качестве аргументов принимается первоначальная цена и возраст автомобиля, и возвращается ориентировочная цена в дилерском центре.

// CalculateResellPrice(1000, 1)
// // => 800

// CalculateResellPrice(1000, 5)
// // => 700

// CalculateResellPrice(1000, 15)
// // => 500
// Обратите внимание, что возвращаемая стоимость равна float64.
