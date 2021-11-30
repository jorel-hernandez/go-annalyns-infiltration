package annalyn

// CanFastAttack can be executed only when the knight is sleeping
func CanFastAttack(knightIsAwake bool) bool {
	//panic("Please implement the CanFastAttack() function")
	var attack bool
	
	if knightIsAwake == false {
		attack = true
	}
	
	return attack
}

// CanSpy can be executed if at least one of the characters is awake
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	//panic("Please implement the CanSpy() function")
	var spy bool
	
	if knightIsAwake == false && archerIsAwake == false && prisonerIsAwake == false {
		spy = false
	} else {
		spy = true
	}
	
	return spy
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	//panic("Please implement the CanSignalPrisoner() function")
	var signal bool
	
	if archerIsAwake == false && prisonerIsAwake == true {
		signal = true
	}
	
	return signal
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	//panic("Please implement the CanFreePrisoner() function")
	
	var free bool
	
	if petDogIsPresent == true && archerIsAwake == false {
		free = true
	} else if petDogIsPresent == false && prisonerIsAwake == true && knightIsAwake == false && archerIsAwake == false {
		free = true
	} else {
		free = false
	}
	
	return free
}
