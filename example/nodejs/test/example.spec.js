const {expect} = require('chai');
const {example} = require('../src/example');

describe('Example', () => {
    it('should say hello', () => {
        expect(example()).to.equal('hello, world!')
    });
});
