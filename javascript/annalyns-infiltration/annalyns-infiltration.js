// проникновение Анналин

export function canExecuteFastAttack(knightIsAwake) {
  return !knightIsAwake;
} // быстрая атака возвожна если рыцарь спит и не успеет одеть доспехи

export function canSpy(knightIsAwake, archerIsAwake, prisonerIsAwake) {
  return knightIsAwake || archerIsAwake || prisonerIsAwake;
} // шпионить можно если хотя бы кто-то не спит

export function canSignalPrisoner(archerIsAwake, prisonerIsAwake) {
  return !archerIsAwake && prisonerIsAwake;
} // звуки птиц понимает лучник, поэтому он должен спать

export function canFreePrisoner(
  knightIsAwake,
  archerIsAwake,
  prisonerIsAwake,
  petDogIsPresent,
) {
  return (
    (petDogIsPresent && !archerIsAwake) ||
    (prisonerIsAwake && !knightIsAwake && !archerIsAwake)
  );
} // освободить подругу !
