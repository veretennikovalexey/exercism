// NeedsLicense determines whether a license is needed to drive a type of vehicle.
// Only "car" and "truck" require a license.

/**
 * Определяет, требуется ли лицензия для управления транспортным средством определенного типа.
 *
 * @param {string} kind - Тип транспортного средства
 * @returns {boolean} - Требуется ли лицензия
 */
export function needsLicense(kind) {
  return kind === "car" || kind === "truck";
}

// ChooseVehicle recommends a vehicle for selection.
// It always recommends the vehicle that comes first in lexicographical order.

/**
 * Помогает выбрать между двумя вариантами, рекомендуя тот, который идет первым в алфавитном порядке.
 *
 * @param {string} option1 - Первый вариант
 * @param {string} option2 - Второй вариант
 * @returns {string} - Сообщение с рекомендацией
 */
export function chooseVehicle(option1, option2) {
  return option1 < option2
    ? `${option1} is clearly the better choice.`
    : `${option2} is clearly the better choice.`;
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.

/**
 * Рассчитывает ориентировочную цену перепродажи транспортного средства в зависимости от его возраста.
 *
 * @param {number} originalPrice - Первоначальная цена
 * @param {number} age - Возраст транспортного средства
 * @returns {number} - Ожидаемая цена перепродажи
 */
export function calculateResellPrice(originalPrice, age) {
  if (age < 3) {
    return originalPrice * 0.8;
  } else if (age > 10) {
    return originalPrice * 0.5;
  } else {
    return originalPrice * 0.7;
  }
}
