import { PUBLIC_API_URL } from '$env/static/public'

export async function pickOne(numbers, max) {
  const response = await fetch(PUBLIC_API_URL + '/lotteries/pick-one?numbers=' + numbers + '&max=' + max);
  return await response.json();
}