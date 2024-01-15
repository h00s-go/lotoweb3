import { lotteries } from '../lotteries.js';

export async function load({ fetch }) {
  const numbers = await lotteries(fetch, 6, 45);

  return {
    numbers: numbers,
  };
}