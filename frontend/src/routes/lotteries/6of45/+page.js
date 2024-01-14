import { PUBLIC_API_URL } from '$env/static/public'

export async function load({ fetch, params }) {
  const response = await fetch(PUBLIC_API_URL + '/lotteries/pick-one?numbers=6&max=45');
  const numbers = await response.json();

  return {
    numbers: numbers.join(', '),
  };
}