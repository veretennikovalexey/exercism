// main.js
window.onload = function () {
  // Количество лайков и дизлайков
  let likes = 1466;
  let dislikes = 1203;

  // Общее количество голосов
  let totalVotes = likes + dislikes;

  // Вычисляем процент дизлайков
  let dislikePercentage = (dislikes / totalVotes) * 100;

  // Округляем до двух знаков после запятой
  dislikePercentage = dislikePercentage.toFixed(2);

  // Находим элемент по ID и вставляем результат
  document.getElementById("result").innerText =
    "Процент дизлайков: " + dislikePercentage + "%";
};
