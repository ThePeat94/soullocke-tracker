<script lang="ts">

	import type { VariantType } from '../types.ts';
	import type { PropsWithChildren } from '../../util/types.ts';

	type Props = {
		variant : VariantType,
		onClick : () => void,
	};

	let {
		variant,
		onClick,
		children,
	} : PropsWithChildren<Props> = $props()

</script>

<button type="button" class={`btn preset-${variant}-primary-500`} onclick={onClick}>{@render children()}</button>
