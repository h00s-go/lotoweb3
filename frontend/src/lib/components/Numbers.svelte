<script>
  import { PUBLIC_API_URL } from '$env/static/public'
  import { onMount } from 'svelte';

  onMount(async (numbers, max) => {
    const response = await fetch(PUBLIC_API_URL + '/lotteries/pick-one?numbers=' + numbers + '&max=' + max);
    pick = await response.json();
    pick = pick.join(', ');
  });

  export let numbers;
  export let max;
  let pick = 'Loading...';
</script>

<div class="card">
  <div class="card-header">
    <b>{ numbers } / { max }</b>
  </div>
  <div class="card-body">
    {pick}
  </div>
</div>