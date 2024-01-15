import { lotteries } from '../lotteries.js';

export async function load({ fetch }) {
  const numbers = await lotteries(fetch, 7, 35);

  return {
    numbers: numbers,
  };
}