package utils

func CheckUniqueAPIIDs(apiIDs []string, comparisonID string) (bool, []string) {
	apiIDMap := make(map[string]struct{})
	var duplicates []string

	for _, apiID := range apiIDs {
		if _, exists := apiIDMap[apiID]; exists {
			duplicates = append(duplicates, apiID)
		} else {
			apiIDMap[apiID] = struct{}{}
		}
	}

	// Check if the comparisonID is in the list of API IDs
	if _, exists := apiIDMap[comparisonID]; exists {
		return false, duplicates
	} else {
		return true, duplicates
	}
}
func ExtractAPIIDs() []string {
	apis, err := FindAllAPIs()
	if err != nil {
		return nil
	}
	var apiIDs []string

	for _, api := range apis {
		apiIDs = append(apiIDs, api)
	}

	return apiIDs
}

func UniqueCheck(comparisonID string) (bool, []string) {
	apiIDs := ExtractAPIIDs()
	return CheckUniqueAPIIDs(apiIDs, comparisonID)
}
