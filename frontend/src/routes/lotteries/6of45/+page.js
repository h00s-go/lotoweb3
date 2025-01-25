import { lotteries } from '../lotteries.js';

export async function load({ fetch }) {
  const lottery = await lotteries(fetch, 6, 45);

  return {
    lottery,
  };
}