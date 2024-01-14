import { pickOne } from '$lib/lotteries'

export async function load({ fetch, params }) {
  const numbers = await pickOne(6, 45);

  return {
    numbers: numbers.join(', '),
  };
}