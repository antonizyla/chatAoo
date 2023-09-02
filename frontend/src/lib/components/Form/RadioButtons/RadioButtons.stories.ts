import type { Meta, StoryObj } from '@storybook/svelte';

import RadioButtons from './RadioButtons.svelte';

const meta = {
	title: 'Form/RadioButtons',
	component: RadioButtons,
	tags: ['autodocs'],
	argTypes: {
		labels: { control: { type: 'array' } }
	}
} satisfies Meta<RadioButtons>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Primary: Story = {
	args: {
		labels: ['one', 'two', 'three']
	}
};
