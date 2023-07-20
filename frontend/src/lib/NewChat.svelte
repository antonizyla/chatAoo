<script lang="ts">
	let value: string | null;

	function send() {
		if (value) {
			console.log(value);
			value = null;
		}
	}
</script>

<form action="">
	<input type="text" name="newChat" placeholder="New Chat ..." bind:value />
	<button on:click={send}>Send</button>
</form>
