import { fetchLotteries } from '$svc/lotteries';

export async function load({ fetch }) {
  const lottery = await fetchLotteries(fetch, 6, 45);

  return {
    lottery,
  };
}