<script>
  import Lottery from '$comp/common/Lottery.svelte';

  const { data } = $props();
</script>

<Lottery lottery={ data.lottery } />