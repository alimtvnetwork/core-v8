package corestr

type HashmapDataModel struct {
	Items map[string]string `json:"Hashmap"`
}

func NewHashmapUsingDataModel(dataModel *HashmapDataModel) *Hashmap {
	return &Hashmap{
		items:         dataModel.Items,
		hasMapUpdated: false,
		cachedList:    nil,
	}
}

func NewHashmapsDataModelUsing(collection *Hashmap) *HashmapDataModel {
	return &HashmapDataModel{
		Items: collection.items,
	}
}
