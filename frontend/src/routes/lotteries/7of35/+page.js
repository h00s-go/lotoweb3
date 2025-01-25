import { fetchLotteries } from '$svc/lotteries';

export async function load({ fetch }) {
  const lottery = await fetchLotteries(fetch, 7, 35);

  return {
    lottery,
  };
}