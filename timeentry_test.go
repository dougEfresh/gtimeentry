package gtimeentry

import (
	"github.com/dougEfresh/toggl-test"
	"testing"
	"time"
)

func togglClient(t *testing.T) *TimeEntryClient {
	tu := &gtest.TestUtil{}
	client := tu.MockClient(t)
	return NewClient(client)
}

func TestTimeEntryDelete(t *testing.T) {
	tClient := togglClient(t)
	err := tClient.Delete(1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTimeEntryList(t *testing.T) {
	tClient := togglClient(t)
	te, err := tClient.List()
	if err != nil {
		t.Fatal(err)
	}
	if len(te) < 1 {
		t.Fatal("<1")
	}

}

func TestTimeEntryCreate(t *testing.T) {
	tClient := togglClient(t)

	te := &TimeEntry{}
	te.Billable = false
	te.Duration = 1200
	te.Pid = 123
	te.Wid = 777
	te.Description = "Meeting with possible clients"
	te.Tags = []string{"billed"}

	nTe, err := tClient.Create(te)

	if err != nil {
		t.Fatal(err)
	}
	if nTe.Id != 3 {
		t.Error("!= 3")
	}
}

func TestTimeEntryUpdate(t *testing.T) {
	tClient := togglClient(t)
	te, err := tClient.Get(1)
	if err != nil {
		t.Fatal(err)
	}
	te.Description = "new"
	nTe, err := tClient.Update(te)
	if err != nil {
		t.Fatal(err)
	}
	if nTe.Description != "new" {
		t.Error("!= new")
	}
}

func TestTimeEntryGet(t *testing.T) {
	tClient := togglClient(t)

	timeentry, err := tClient.Get(1)
	if err != nil {
		t.Fatal(err)
	}
	if timeentry.Id != 1 {
		t.Error("!= 1")
	}

	st, err := time.Parse(time.RFC3339, "2013-02-27T01:24:00+00:00")

	if err != nil {
		t.Fatal(err)
	}
	diff := st.Sub(timeentry.Start)
	if diff != 0 {
		t.Errorf("!= %s", diff)
	}
	st, err = time.Parse(time.RFC3339, "2013-02-27T07:24:00+00:00")
	diff = st.Sub(timeentry.Stop)
	if diff != 0 {
		t.Errorf("!= %s", diff)
	}

	/*
		if timeentry.FullName != "John Swift" {
			t.Error("!= John Swift:  " + timeentry.FullName)
		}

		if timeentry.ApiToken != "1971800d4d82861d8f2c1651fea4d212" {
			t.Error("!= J1971800d4d82861d8f2c1651fea4d212:  " + timeentry.ApiToken)
		}

		if timeentry.Email != "johnt@swift.com" {
			t.Error("!= johnt@swift.com" + timeentry.Email)
		}
	*/
}
