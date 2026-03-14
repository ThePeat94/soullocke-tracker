<script lang="ts">

	import type { Snippet } from 'svelte';

	type Props = {
		header: Snippet;
		content: Snippet;
		footer?: Snippet;
		onclick?: () => void;
	}

	const {
		header,
		content,
		footer,
		onclick,
	} : Props = $props();

</script>

<div
	class="card preset-filled-surface-100-900 border border-surface-200-800 card-hover divide-surface-200-800 block divide-y overflow-hidden"
	role="cell"
	tabindex={1}
	onclick={onclick}
	onkeydown={(k) => {
		if (k.code === 'Enter') {
			onclick?.();
		}
	}}
>
	<header class="p-4">
		{@render header()}
	</header>
	<section class="p-4">
		{@render content()}
	</section>
	{#if footer}
		<footer class="p-4">
			{@render footer()}
		</footer>
	{/if}
</div>
