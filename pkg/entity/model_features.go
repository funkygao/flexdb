package entity

func (f *ModelFeature) EnableCRUD() {
	f.EnableCreateRow(true)
	f.EnableReadRow(true)
	f.EnableUpdateRow(true)
	f.EnableDeleteRow(true)
}

func (f *ModelFeature) EnableChangeAudit(yes bool) {
	if yes {
		f[0] = '1'
	} else {
		f[0] = '0'
	}
}

func (f *ModelFeature) ChangeAuditEnabled() bool {
	return f[0] == '1'
}

func (f *ModelFeature) EnableFakeDelete(yes bool) {
	if yes {
		f[1] = '1'
	} else {
		f[1] = '0'
	}
}

func (f *ModelFeature) FakeDeleteEnabled() bool {
	return f[1] == '1'
}

func (f *ModelFeature) EnableComment(yes bool) {
	if yes {
		f[2] = '1'
	} else {
		f[2] = '0'
	}
}

func (f *ModelFeature) CommentEnabled() bool {
	return f[2] == '1'
}

func (f *ModelFeature) EnableReadRow(yes bool) {
	if yes {
		f[3] = '1'
	} else {
		f[3] = '0'
	}
}

func (f *ModelFeature) ReadRowEnabled() bool {
	return f[3] == '1'
}

func (f *ModelFeature) EnableUpdateRow(yes bool) {
	if yes {
		f[4] = '1'
	} else {
		f[4] = '0'
	}
}

func (f *ModelFeature) UpdateRowEnabled() bool {
	return f[4] == '1'
}

func (f *ModelFeature) EnableCreateRow(yes bool) {
	if yes {
		f[5] = '1'
	} else {
		f[5] = '0'
	}
}

func (f *ModelFeature) CreateRowEnabled() bool {
	return f[5] == '1'
}

func (f *ModelFeature) EnableDeleteRow(yes bool) {
	if yes {
		f[6] = '1'
	} else {
		f[6] = '0'
	}
}

func (f *ModelFeature) DeleteRowEnabled() bool {
	return f[6] == '1'
}
