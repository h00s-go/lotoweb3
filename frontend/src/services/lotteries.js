import { PUBLIC_API_URL } from '$env/static/public'

export async function fetchLotteries(fetch, numbers, max) {
  const response = await fetch(PUBLIC_API_URL + '/lotteries/pick-many?count=10&numbers=' + numbers + '&max=' + max);
  let result = await response.json();

  return {
    numbers,
    max,
    picks: result,
  };
}