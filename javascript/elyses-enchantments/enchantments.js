export function getItem(cards, position) {
  return cards[ position ];
}

export function setItem(cards, position, replacementCard) {
  cards[ position ] = replacementCard;
  return cards;
}

export function insertItemAtTop(cards, newCard) {
  cards.push( newCard );
  return cards;
}

export function removeItem(cards, position) {
  cards.splice(position, 1);                             // Удаляем один элемент по заданной позиции
  return cards;                                          // Возвращаем изменённый массив
}

export function removeItemFromTop(cards) {
  cards.pop();
  return cards;                                          // Возвращаем изменённый массив
}

export function insertItemAtBottom(cards, newCard) {
  cards.unshift( newCard );
  return cards;                                          // Возвращаем изменённый массив
}

export function removeItemAtBottom(cards) {
  cards.shift();
  return cards;                                          // Возвращаем изменённый массив
}

export function checkSizeOfStack(cards, stackSize) {
  return cards.length === stackSize;                     // Сравниваем длину массива с заданным размером
}
