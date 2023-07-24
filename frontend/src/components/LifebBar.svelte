<script>
    let dragon = "🐲";
    let hero = "😃";
    let distance;
    export let lives;
    export let keyChangeCounter;

    $: {
        distance = lives - keyChangeCounter >= 0 ? lives - keyChangeCounter : 0;
        if (distance/lives > 0.7) {
            hero = "😃";
        } else if (distance/lives > 0.5) {
            hero = "😐";
        } else if (distance/lives > 0.25) {
            hero = "😨";
        } else if (distance > 0) {
            hero = "😱";
        } else {
            hero = "💀";
        }
    }
</script>

<div class="life-bar">
    <h1 class="title is-1">
        <span>{hero}</span>
        <span class="life-bar__distance">
            {#each Array(distance) as _, i (i)}
            <span>_</span>
            {/each}
        </span>
        <span>{dragon}</span>
    </h1>
</div>

<style>

</style>
