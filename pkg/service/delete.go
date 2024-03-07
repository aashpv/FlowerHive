package service

func (s *service) Delete(id int) (err error) {
	err = s.pgs.Delete(id)
	if err != nil { // if err use log for writing in file
		return
	}
	// если ошибок нет, то тоже надо логировать

	return
}
