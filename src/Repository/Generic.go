package repository

func FindAll(entities interface{}) error {
	result := DB.Find(entities)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func Update(entity interface{}) error {
	result := DB.Save(entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindByID(id uint, entity interface{}) error {
	result := DB.First(entity, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func Create(entity interface{}) error {
	result := DB.Create(entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func Delete(entity interface{}) error {
	result := DB.Delete(entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
