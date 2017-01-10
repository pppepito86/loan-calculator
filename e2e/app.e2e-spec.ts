import { LoanCalculatorPage } from './app.po';

describe('loan-calculator App', function() {
  let page: LoanCalculatorPage;

  beforeEach(() => {
    page = new LoanCalculatorPage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});
