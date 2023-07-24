<script>
    import { onMount } from "svelte";
    import { writable, derived } from "svelte/store";
    import LifeBar from './components/LifebBar.svelte';

    let game = writable({});
    let key = writable("");
    let success = writable(false);
    let error = writable("");
    let apiUrl = "http://localhost:8000";
    let prevKey = writable(""); 
    export let keyChangeCounter = writable(0); 
    export const lives = 30;

    const processedChallenge = derived([key, game], ([$key, $game]) => {
        if ($game.challenge) {
            let tempChallenge = $game.challenge.split("");
            for (let i = 0; i < $key.length; i++) {
                let char = $key[i];
                tempChallenge = tempChallenge.map((ch) =>
                    ch == (i + 1).toString() ? char : ch
                );
            }
            return tempChallenge.join("");
        } else {
            return "Loading game...";
        }
    });

    let keyFull = derived(
        [key, game, success],
        ([$key, $game, $success]) =>
            $key.length === $game.keyLength && !$success
    );

    const checkKey = async () => {
        if ($key.length !== ($game.keyLength || $success)) {
            return;
        }

        try {
            const res = await fetch(`${apiUrl}/api/v1/key`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    gameId: $game.id,
                    key: $key,
                }),
            });
            if (!res.ok) {
                throw new Error("Failed to check the key");
            }
            const data = await res.json();
            success.set(data.success);
        } catch (e) {
            error.set(e.message);
            console.error(e);
        }
    };

    key.subscribe(($key) => {
        let $prevKey;
        prevKey.subscribe(value => $prevKey = value)(); // get the value of prevKey
        prevKey.set($key); // set prevKey as current key

        if ($prevKey.length <= $key.length) { // check if a character has been added
            for (let i = 0; i < $key.length; i++) {
                if ($key[i] !== $prevKey[i]) { // check if a character has changed
                    keyChangeCounter.update(n => n + 1);
                }
            }
        }
        checkKey();
    });

    onMount(async () => {
        try {
            const res = await fetch(`${apiUrl}/api/v1/game`);
            if (!res.ok) {
                throw new Error("Failed to fetch the game");
            }
            game.set(await res.json());
        } catch (e) {
            error.set(e.message);
            console.error(e);
        }
    });

    const onInput = (event) => {
        let newValue = event.target.value;

        // Key length
        if (newValue.length > $game.keyLength) {
            newValue = newValue.slice(0, $game.keyLength);
        }

        // Only letter
        newValue = newValue.replace(/[^a-zA-Z]/g, "");

        key.set(newValue);
    };
</script>

<main>
    <div class="container">
        <section class="section">
            <h1 class="title is-1 has-text-centered">Cryptle</h1>

            <div class="columns is-centered">
                {#each Array($game.keyLength).fill(0) as _, index (index)}
                    <div class="column is-1">
                        <div
                            class={`key-cell ${
                                $key[index] ? "has-background-info" : ""
                            } ${
                                $success
                                    ? "has-background-success"
                                    : $key.length === $game.keyLength
                                    ? "has-background-danger"
                                    : ""
                            }`}
                        >
                            <p>{index + 1}</p>
                            <p>{$key[index] || ""}</p>
                        </div>
                    </div>
                {/each}
            </div>
        </section>

        <h2 class="title is-2 has-text-centered">{$keyChangeCounter}</h2>
       
        <LifeBar lives={lives} keyChangeCounter={$keyChangeCounter} />
        
        <section class="section">
            <div class="content is-large">
                <p>{$processedChallenge}</p>
            </div>
        </section>

        <section class="section">
            <div class="control">
                <input
                    class={`input ${$success ? "is-success" : ""} ${
                        $keyFull ? "is-danger" : ""
                    }`}
                    type="text"
                    bind:value={$key}
                    on:input={onInput}
                    placeholder="Enter your key..."
                    disabled={$success}
                />
            </div>
        </section>

        {#if $success}
            <p>Success!</p>
        {:else if $error}
            <p>{$error}</p>
        {:else}
            <p />
        {/if}
    </div>

    <footer class="footer is-widescreen">
        <div class="content has-text-centered">
            <h5 class="title is-6 has-text-centered has-text-grey-light">
                Game ID: {$game.id}
            </h5>
        </div>
    </footer>
</main>

<style global lang="sass" >
    $quantity: 15
  
    .key-cells 
      display: flex
  
    .key-cell 
      width: 60px
      height: 60px
      border: 1px solid #000
      border-radius: 8px
      display: flex
      flex-direction: column
      align-items: center
      justify-content: center
      margin-right: 5px
      transition: background-color 0.5s ease
  
    .firefly
      position: fixed
      left: 50%
      top: 50%
      width: 0.4vw
      height: 0.4vw
      margin: -0.2vw 0 0 9.8vw
      animation: ease 200s alternate infinite
      pointer-events: none
  
      &::before,
      &::after
        content: ''
        position: absolute
        width: 100%
        height: 100%
        border-radius: 50%
        transform-origin: -10vw
  
      &::before
        background: black
        opacity: 0.4
        animation: drift ease alternate infinite
  
      &::after
        background: white
        opacity: 0
        box-shadow: 0 0 0vw 0vw yellow
        animation: drift ease alternate infinite, flash ease infinite
  
    @for $i from 1 through $quantity
      $steps: random(12) + 16
      $rotationSpeed: random(10) + 8s
  
      .firefly:nth-child(#{$i})
        animation-name: move#{$i}
  
        &::before
          animation-duration: #{$rotationSpeed}
  
        &::after
          animation-duration: #{$rotationSpeed}, random(6000) + 5000ms
          animation-delay: 0ms, random(8000) + 500ms
  
      @keyframes move#{$i}
        @for $step from 0 through $steps
          #{$step * (100 / $steps)}%
            transform: translateX(random(100) - 50vw) translateY(random(100) - 50vh) scale(random(75) / 100 + .25)
  
    @keyframes drift
      0%
        transform: rotate(0deg)
      100%
        transform: rotate(360deg)
  
    @keyframes flash
      0%, 30%, 100%
        opacity: 0
        box-shadow: 0 0 0vw 0vw yellow
      5%
        opacity: 1
        box-shadow: 0 0 2vw 0.4vw yellow
</style>
