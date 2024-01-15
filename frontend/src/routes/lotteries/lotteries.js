import { PUBLIC_API_URL } from '$env/static/public'

export async function lotteries(fetch, numbers, max) {
  const response = await fetch(PUBLIC_API_URL + '/lotteries/pick-one?numbers=' + numbers + '&max=' + max);
  let result = await response.json();
  return {
    numbers: numbers,
    max: max,
    pick: result.join(', '),
  };
}