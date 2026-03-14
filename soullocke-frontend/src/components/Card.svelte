<script lang="ts">
	export let onClick: () => void = () => {};
</script>

<div
	class="card preset-filled-surface-100-900 border border-surface-200-800 card-hover divide-surface-200-800 block divide-y overflow-hidden"
	role="cell"
	tabindex={1}
	on:click={onClick}
	on:keydown={(k) => {
		if (k.code === 'Enter') {
			onClick();
		}
	}}
>
	<header class="p-4">
		<slot name="header" />
	</header>
	<section class="p-4">
		<slot name="content" />
	</section>
	{#if $$slots.footer}
		<footer class="p-4 grid grid-cols-2">
			<slot name="footer" />
		</footer>
	{/if}
	{#if $$slots.custom}
		<section class="p-4">
			<slot name="custom" />
		</section>
	{/if}
</div>
