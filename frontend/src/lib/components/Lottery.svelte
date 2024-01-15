<script>
  export let lottery;
</script>

<div class="card">
  <div class="card-header">
    <b>{ lottery.numbers } / { lottery.max }</b>
  </div>
  <div class="card-body">
    <p>{ lottery.pick }</p>
  </div>
</div>