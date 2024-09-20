// Лазанья
package lasagna

// Время выпечки
const OvenTime = 40 

// Оставшееся время выпечки
func RemainingOvenTime(actualMinutesInOven int) int {
    return OvenTime - actualMinutesInOven
}

// Для каждого слоя лазаньи требуется 2 минуты
func PreparationTime(numberOfLayers int) int {
	return numberOfLayers + numberOfLayers
} // each layer takes you 2 minutes to prepare

// Тут я не совсем понял
// но это сумма времени приготовления PreparationTime + время в печи actualMinutesInOven
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return numberOfLayers + numberOfLayers + actualMinutesInOven
}

// numberOfLayers      число слоёв
// actualMinutesInOven в печи