import { lotteries } from '../lotteries.js';

export async function load({ fetch }) {
  const lottery = await lotteries(fetch, 7, 35);

  return {
    lottery: lottery,
  };
}