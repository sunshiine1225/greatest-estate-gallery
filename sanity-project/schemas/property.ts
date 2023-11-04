import {defineType} from 'sanity'

export default {
	name: 'property',
	type: 'document',
	title: 'Property',
	fields: [
		defineType({
			title: 'Property Name',
			name: 'property_name',
			type: 'string',
		}),
		defineType({
			title: 'Images',
			name: 'images',
			type: 'array',
			of: [
				{
					type: 'image',
				},
			],
		}),
		defineType({
			type: 'object',
			name: 'price',
			title: 'Price',
			fields: [
				{type: 'number', name: 'price', title: 'Price'},
				{
					type: 'string',
					name: 'denomination',
					title: 'Denomination',
					options: {
						list: ['Cr', 'L'],
					},
				},
			],
		}),
		defineType({
			title: 'Num Bedrooms',
			name: 'num_bedrooms',
			type: 'number',
		}),
		defineType({
			title: 'Num Bathrooms',
			name: 'num_bathrooms',
			type: 'number',
		}),
		defineType({
			title: 'Size (sq. ft)',
			name: 'size',
			type: 'number',
		}),
		defineType({
			title: 'Type',
			name: 'type',
			type: 'string',
			options: {
				list: ['Duplex', 'Apartments'],
			},
		}),
		defineType({
			name: 'location',
			title: 'Location',
			type: 'object',
			fields: [
				{type: 'geopoint', name: 'geolocation', title: 'GeoLocation'},
				{type: 'string', name: 'area', title: 'Area'},
				{type: 'string', name: 'city', title: 'City'},
			],
		}),
	],
}
