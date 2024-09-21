// Главная героиня игры - Анналин, храбрая девушка со свирепой и преданной собакой

package annalyn

func CanFastAttack(knightIsAwake bool) bool {
	return !knightIsAwake
} // CanFastAttack can be executed only when the knight is sleeping.

func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	return knightIsAwake || archerIsAwake || prisonerIsAwake
} // CanSpy can be executed if at least one of the characters is awake.

func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
    return !archerIsAwake && prisonerIsAwake
} // CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.

func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
    if petDogIsPresent && !archerIsAwake {
        return true
    } else if prisonerIsAwake && !knightIsAwake && !archerIsAwake {
        return true         
    } else {
        return false
    }
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
