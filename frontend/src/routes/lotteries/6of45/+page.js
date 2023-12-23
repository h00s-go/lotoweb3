export async function load({ params }) {
  const response = await fetch('http://localhost:3000/api/v1/lotteries/6of45');
  const numbers = await response.json();

  return {
    numbers: numbers.join(', '),
  };
}