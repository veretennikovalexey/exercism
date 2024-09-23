package cards

// FavoriteCards возвращает список любимых карт
func FavoriteCards() []int {
  return []int{ 2, 6, 9 }
}

func GetItem(slice []int, index int) int {
  if index >= 0 && index < len(slice) {
    return slice[index]
  }
  return -1
}

func SetItem(slice []int, index, value int) []int {
    if index < 0 || GetItem( slice, index ) < 0 {
        return append( slice, value )
    } else {
        slice[ index ] = value
        return slice
    }
}

// PrependItems добавляет произвольное количество значений в начало среза.
func PrependItems(slice []int, values ...int) []int {
	return append(values, slice...)
}

// RemoveItem удаляет элемент из среза по указанному индексу. Если индекс вне диапазона, возвращает исходный срез.
func RemoveItem(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}
