import type { Meta, StoryObj } from '@storybook/svelte';

import Image from './Image.svelte';

const meta: Meta = {
	title: 'Image',
	component: Image,
	argTypes: {
		alt: {
			control: 'text'
		},
		title: {
			control: 'text'
		},
		description: {
			control: 'text'
		},
		pad: {
			control: 'boolean'
		}
	}
} satisfies Meta<Image>;

export default meta;
type Story = StoryObj<typeof Image>;

import sImage from './sampleImage.jpg?format=webp&width=400;800&picture';

export const Primary: Story = {
	args: {
		data: sImage,
		alt: 'Sample Image',
		title: 'Sample Image',
		description: 'This is a sample image',
		pad: true
	}
};
