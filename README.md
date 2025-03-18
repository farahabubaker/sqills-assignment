# sqills-assignment

For this assignment, I have implemnted a ticketing service that handles incoming reservation requests.

The test cases I have implemnted follow the 4 scenarios mentioned: 

1. Be able to make a reservation for 2 passengers with 2 first-class seats, from Paris to
Amsterdam on service 5160 on April 1st 2021 with seats A11 & A12 (Carriage A, Seat
11 & 12).
   
2. Make the same booking again, but then it should fail, because seats are already taken.
   
3. Be able to make a reservation for 2 passengers where 1 passenger in second-class, one
in first-class. From London to Amsterdam. London to Paris, seat H1 and N5, Paris to
Amsterdam, seat A1 & T7.

4. Make the previous booking again, and it should fail, because seats are taken.

## How to run the tests
### Integration Tests
In main_test.go, there are two test cases that follow the scenarios mentioned above.

Test_IntegrationScenario1()
- This test case runs scenario 1 & 2
- go test -run Test_IntegrationScenario1 -v

Test_IntegrationScenario2()
- This test case runs scenario 3 & 4
- go test -run Test_IntegrationScenario2 -v
 
The tests will first validate a reservation with the given passenger information, route information, etc., if all is valid, then the booking is created.
The test will run again with the same mock request and fail at validation since the seats are no longer available

### Business Logic tests
In handlers/reservationHandler_test.go, there are two test cases that test the ValidateReservation handler and the CreateReservation handler.

### Database Togic tests
In repositories/booking_test.go, there are two tests cases that test the VerifyBooking database call and the Create booking database call.

