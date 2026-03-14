<script lang="ts">
	type Props = {
		value: string
		label: string
		placeholder?: string
	};

	let {
		value = $bindable<string>(),
		label,
		placeholder,
	} : Props = $props();
</script>


<label class="label">
	<span class="label-text">{label}</span>
	<input type="text" class="input p-3 rounded-md border hover:border-gray-50" bind:value={value} placeholder={placeholder}/>
</label>
