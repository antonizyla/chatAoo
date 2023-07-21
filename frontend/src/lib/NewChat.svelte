<script lang="ts">
	let value: string | null;

	function send() {
		if (value) {
			value = null;
		}
	}
</script>

<form action="">
	<input type="text" name="newChat" placeholder="New Chat ..." bind:value />
	<button on:click={send}>Send</button>
</form>
